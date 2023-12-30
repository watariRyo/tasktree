// Import the functions you need from the SDKs you need
import { getApps, initializeApp } from 'firebase/app';
import { getAnalytics } from 'firebase/analytics';
import { browser } from '$app/environment';
import {
	PUBLIC_API_ID,
	PUBLIC_API_KEY,
	PUBLIC_AUTH_DOMAIN,
	PUBLIC_MEASUREMENT_ID,
	PUBLIC_MESSAGING_SENDER_ID,
	PUBLIC_PROJECT_ID,
	PUBLIC_STORAGE_BUCKET
} from '$env/static/public';
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
	apiKey: PUBLIC_API_KEY,
	authDomain: PUBLIC_AUTH_DOMAIN,
	projectId: PUBLIC_PROJECT_ID,
	storageBucket: PUBLIC_STORAGE_BUCKET,
	messagingSenderId: PUBLIC_MESSAGING_SENDER_ID,
	appId: PUBLIC_API_ID,
	measurementId: PUBLIC_MEASUREMENT_ID
};

// Initialize Firebase
if (getApps().length === 0) {
	console.log('firebase initialization');
	const app = initializeApp(firebaseConfig);
	if (browser) {
		getAnalytics(app);
	}
}
