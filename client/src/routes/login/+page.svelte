<script lang="ts">
	import { page } from '$app/stores';
	import AuthForm from '$lib/components/Auth/AuthForm.svelte';
	import LoginWithGoogle from '$lib/components/Auth/LoginWithGoogle.svelte';
	import { loginWithEmailAndPassowrd } from '$lib/firebase/auth.client';
	import { afterLogin } from '$lib/helpers/route.helper';
	import messagesStore from '$lib/stores/messages.store';

	const onLogin = async (e: SubmitEvent) => {
		try {
			const formData = new FormData(e.target as HTMLFormElement);
			const email = formData.get('email')?.toString() || '';
			const password = formData.get('password')?.toString() || '';
			if (email === '' || password === '') {
                messagesStore.showError('Please enter a valid email and password.');
				return;
			}
			const user = await loginWithEmailAndPassowrd(email, password);
            messagesStore.hide()
			await afterLogin($page.url, user.uid);
		} catch (error: any) {
			// console.log(error);
            messagesStore.showError();
		}
	};
</script>

<div class="row">
    <div class="col">
        <h1>Login</h1>
    </div>
</div>
<AuthForm on:submit={onLogin} btnName="Login" />
<hr />
<LoginWithGoogle isSignup={false} />
<hr />
<div class="row">
	<div class="col">
		<a href="/forgot-password" class="btn btn-warning w-100">Forgot Password</a>
	</div>
</div>

<svelte:head>
	<title>Book Lovers - Login</title>
</svelte:head>