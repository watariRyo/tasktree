<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import AuthForm from '$lib/components/Auth/AuthForm.svelte';
	import LoginWithGoogle from '$lib/components/Auth/LoginWithGoogle.svelte';
	import { registerWithEmailAndPassword } from '$lib/firebase/auth.client';
	import { afterRegist } from '$lib/helpers/route.helper';
	import messagesStore from '$lib/stores/messages.store';

    const register = async (e: SubmitEvent) => {
		try {
			const formData = new FormData(e.target as HTMLFormElement);
			const email = formData.get('email')?.toString() || '';
			const password = formData.get('password')?.toString() || '';
			if (email === '' || password === '') {
                messagesStore.showError('Please enter a valid email and password.');
				return;
			}
			if (password.length < 8) {
                messagesStore.showError('Password must be at least 8 characters.');
				return;
			}
			const user = await registerWithEmailAndPassword(email, password);
            messagesStore.hide()
			await afterRegist($page.url);
		} catch (error: any) {
			if (error.code === 'auth/email-already-in-use') {
                messagesStore.showError('You have already registered, please log in.');
				await goto('/login');
			} else {
                messagesStore.showError();
			}
		}
	};
</script>

<div class="row">
	<div class="col">
		<h1>Sign Up</h1>
	</div>
</div>
<hr />
<AuthForm on:submit={register} btnName="Sign Up" />
<hr />
<LoginWithGoogle isSignup={true} />

<svelte:head>
	<title>TaskTree - Sign Up</title>
</svelte:head>
