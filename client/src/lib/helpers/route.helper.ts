import { goto } from '$app/navigation';
import { sendJWTToken } from '$lib/firebase/auth.client';

export const afterRegist = async (url: URL, userId: string) => {
	const route = url.searchParams.get('redirect') || '/';
	// await setUser(userId);
	console.log(userId);
	await sendJWTToken();
	await goto(route);
};

export const afterLogin = async (url: URL, userId: string) => {
	const route = url.searchParams.get('redirect') || '/';
	console.log(userId);
	await sendJWTToken();
	await goto(route);
};
