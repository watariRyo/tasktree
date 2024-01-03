import { goto } from '$app/navigation';
import { sendJWTToken } from '$lib/firebase/auth.client';

export const afterRegist = async (url: URL) => {
	const route = url.searchParams.get('redirect') || '/';
	// await setUser(userId);
	await sendJWTToken();
	await goto(route);
};

export const afterLogin = async (url: URL) => {
	const route = url.searchParams.get('redirect') || '/';
	await sendJWTToken();
	await goto(route);
};
