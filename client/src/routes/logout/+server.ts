import { json, type RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async ({ cookies }) => {
	cookies.delete('jwt', { path: '/' });
	cookies.delete('ssid', { path: '/' });
	return json({ message: 'success' }, { status: 200 });
};
