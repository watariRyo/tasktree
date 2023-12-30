import { auth } from '$lib/firebase/firebase.server';
import { redirect, type Handle } from '@sveltejs/kit';
import type { User } from './models/user';

export const handle: Handle = async ({ event, resolve }) => {
	const protectRoutes = ['/add', '/edit', '/profile'];

	const guestRoutes = ['/login', '/signup', '/forgot-password'];

	try {
		const user = await getFirebaseUser(event.cookies.get('jwt'));
		event.locals.user = user;
	} catch (e: any) {
		event.locals.user = null;
	}

	const user = event.locals?.user;
	const url = event.url;

	if (url.pathname !== '/') {
		if (!user && protectRoutes.find((p) => url.pathname.indexOf(p) > -1)) {
			throw redirect(302, `/login?redirect=${url.pathname}`);
		}
		if (user && guestRoutes.find((u) => url.pathname.indexOf(u) > -1)) {
			throw redirect(302, '/');
		}
	}

	const response = await resolve(event);
	return response;
};

const getFirebaseUser = async (token: string | undefined) => {
	if (!token) {
		return null;
	}
	const decodedToken = await auth.verifyIdToken(token, true);
	const authUser = await auth.getUser(decodedToken.uid);

	const email = authUser.email || '';
	const user: User = { id: authUser.uid, email: email };

	return user;
};
