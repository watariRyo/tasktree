import { PUBLIC_BACKEND_FQDN } from '$env/static/public';
import { json, type RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async ({ cookies, fetch }) => {
	await fetch(`${PUBLIC_BACKEND_FQDN}/api/v1/logout`, {
		method: 'GET'
	});
	cookies.delete('jwt', { path: '/' });
	cookies.delete('ssid', { path: '/' });
	return json({ message: 'success' }, { status: 200 });
};
