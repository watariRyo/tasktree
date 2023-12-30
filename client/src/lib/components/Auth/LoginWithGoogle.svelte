<script lang="ts">
	import { page } from '$app/stores';
	import { loginWithGoogle } from '$lib/firebase/auth.client';
	import { afterLogin, afterRegist } from '$lib/helpers/route.helper';

	export let isSignup: boolean;

	const loginGoogle = async () => {
		try {
			const user = await loginWithGoogle();
			if (isSignup) {
				await afterRegist($page.url, user.uid);
			} else {
				await afterLogin($page.url, user.uid);
			}
		} catch (e: any) {
			if (e.code === 'auth/popup-closed-by-user') {
				return;
			}
			console.log(e);
		}
	};
</script>

<div class="row">
	<div class="col">
		<button on:click={loginGoogle} class="btn btn-primary w-100">Login With Google</button>
	</div>
</div>
