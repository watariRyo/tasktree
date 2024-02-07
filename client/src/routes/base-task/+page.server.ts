import { PUBLIC_BACKEND_FQDN } from '$env/static/public';

export const load = async ({ fetch }) => {
	try {
		const res = await fetch(`${PUBLIC_BACKEND_FQDN}/api/v1/base`, {
			method: 'GET'
		});
		const data = await res.json();

		return {
			status: data.status,
			code: data.code
		};
		// eslint-disable-next-line @typescript-eslint/no-explicit-any
	} catch (e: any) {
		console.log('error');
	}
};
