import { client } from '$lib/helpers/api.helper';
import {
	createUserWithEmailAndPassword,
	getAuth,
	GoogleAuthProvider,
	sendPasswordResetEmail,
	signInWithEmailAndPassword,
	signInWithPopup,
	signOut,
	type User
} from 'firebase/auth';
import { ApiError } from '../../models/api.error';

export const loginWithGoogle = async (): Promise<User> => {
	const auth = getAuth();
	const userCredential = await signInWithPopup(auth, new GoogleAuthProvider());
	return userCredential.user;
};

export const logout = async () => {
	await signOut(getAuth());
	await fetch('/logout');
};

export const registerWithEmailAndPassword = async (
	email: string,
	password: string
): Promise<User> => {
	const userCredential = await createUserWithEmailAndPassword(getAuth(), email, password);
	return userCredential.user;
};

export const loginWithEmailAndPassowrd = async (email: string, password: string): Promise<User> => {
	const userCredential = await signInWithEmailAndPassword(getAuth(), email, password);
	return userCredential.user;
};

export const mailResetpasswordEmail = async (email: string) => {
	await sendPasswordResetEmail(getAuth(), email);
};

export const sendJWTToken = async () => {
	const auth = getAuth();
	const user = auth.currentUser;

	if (!user) {
		return;
	}

	// const token = await user.getIdToken();
	await client('/token', {
		method: 'POST',
		body: JSON.stringify({
			token:
				'eyJhbGciOiJSUzI1NiIsImtpZCI6IjUyNmM2YTg0YWMwNjcwMDVjZTM0Y2VmZjliM2EyZTA4ZTBkZDliY2MiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoi5rih5ra8IiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0llNURaWFBIcm9kWk9vRUt2VEhCVUNXbnpubHl0TjhPYklHWmZ5M0RVX1pvVT1zOTYtYyIsImlzcyI6Imh0dHBzOi8vc2VjdXJldG9rZW4uZ29vZ2xlLmNvbS90YXNrdHJlZS1iOWRiMCIsImF1ZCI6InRhc2t0cmVlLWI5ZGIwIiwiYXV0aF90aW1lIjoxNzA0MTYwNTk1LCJ1c2VyX2lkIjoiNHlMcEhyRkxIaGFyY2VPOHgyM3dXeVd2NHZNMiIsInN1YiI6IjR5THBIckZMSGhhcmNlTzh4MjN3V3lXdjR2TTIiLCJpYXQiOjE3MDQxNjA1OTYsImV4cCI6MTcwNDE2NDE5NiwiZW1haWwiOiJkLnNpbHYudmMxOTE5QGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7Imdvb2dsZS5jb20iOlsiMTAxMDI3Mzg3MDYyNzA3MjkwNzY4Il0sImVtYWlsIjpbImQuc2lsdi52YzE5MTlAZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoiZ29vZ2xlLmNvbSJ9fQ.mlz1REMMX-q6aV8y6MA6Kgip2OA-DgzzYFEvRiv4adiVYlv_1zZcE8dvoGlEkjuz1qNle1dww615xZqCmbaZVxezLar7JGdQd5tSsGH3fwJrA_NEPuRc_7gkrFrkh_STMkZflKncns04nOIutLvcabTZ6ERGd2KSVwlxA8NZ68NkEnuLEFAbjZVUFArVhSV4QKKZ6vtuGT56O4RKQQQ30zAiOEu7mI9UQIl-ObRVinIiO3QvCyHq_w2WlK8XuD4LgGwyZXrerKdjUNTJ-jLbgQTsM9yKaHs-rsBCJT6BzRgqRn_hgOanpXAMogh8es2WwnYb3IEMGruG_djtXNU7HQ',
			refreshToken: user.refreshToken,
			email: user.email
		})
	}).catch((error: ApiError) => {
		return error.serverErrorContent?.code;
	});
};
