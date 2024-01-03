import { json, type Cookies, type RequestHandler } from '@sveltejs/kit';
import { auth } from '$lib/firebase/firebase.server';
import { PUBLIC_API_KEY } from '$env/static/public';
import { GrantType } from '../../models/grant.type';

const TOKEN_EXPIRED = 'auth/id-token-expired';

/**@type {import ('./$types').RequestHandler} */
export const POST: RequestHandler = async ({ request, cookies }) => {
	try {
		const { token, refreshToken, email, grantType } = await request.json();
		if (grantType === GrantType.REFRESH_TOKEN) {
			const res = await sendRefreshToken(cookies, refreshToken, email);
			return res;
		} else {
			try {
				const verifyIdToken = await auth.verifyIdToken(token ?? '', true);
				setCookie(cookies, token, verifyIdToken.exp, verifyIdToken.email, email);
				return json({ message: 'success' }, { status: 200 });
				// eslint-disable-next-line @typescript-eslint/no-explicit-any
			} catch (error: any) {
				if (error.code === TOKEN_EXPIRED) {
					const res = await sendRefreshToken(cookies, refreshToken, email);
					return res;
					// return json({ message: 'Access denied', code: 'hoge' }, { status: 403 });
				} else {
					return json({ message: 'Access denied' }, { status: 403 });
				}
			}
		}
	} catch (error) {
		return json({ message: 'Access denied' }, { status: 403 });
	}
};

const sendRefreshToken = async (
	cookies: Cookies,
	refreshToken: string,
	email: string
): Promise<Response> => {
	const res = await fetch(`https://securetoken.googleapis.com/v1/token?key=${PUBLIC_API_KEY}`, {
		method: 'POST',
		body: JSON.stringify({
			grant_type: GrantType.REFRESH_TOKEN,
			refresh_token: refreshToken
		})
	});
	if (res.ok) {
		const data = await res.json();
		const verifyIdToken = await auth.verifyIdToken(data.id_token);
		setCookie(cookies, data.id_token, verifyIdToken.exp, verifyIdToken.email, email);
		return json({ message: 'success' }, { status: 200 });
	} else {
		return json({ message: 'Could not refresh token' }, { status: 401 });
	}
};

const setCookie = (
	cookies: Cookies,
	idToken: string,
	exp: number,
	tokenEmail: string | undefined,
	email: string
) => {
	if (tokenEmail === email) {
		cookies.set('jwt', idToken, {
			maxAge: exp - Date.now() / 1000,
			path: '/'
		});
	}
};
