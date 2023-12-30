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

	const token = await user.getIdToken();
	await fetch('/token', {
		method: 'POST',
		body: JSON.stringify({
			token: token,
			email: user.email
		})
	});
};
