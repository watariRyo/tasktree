import admin from 'firebase-admin';
import serviceAccount from './firebase-secrets.server.json';

if (admin.apps.length === 0) {
	admin.initializeApp({
		credential: admin.credential.cert(serviceAccount as admin.ServiceAccount)
	});
}

export const auth = admin.auth();
