<script lang="ts">
	import { page } from '$app/stores';
	import { loginWithGoogle } from '$lib/firebase/auth.client';
	import { afterLogin, afterRegist } from '$lib/helpers/route.helper';
	import messagesStore from '$lib/stores/messages.store';

	export let isSignup: boolean;

	const loginGoogle = async () => {
		try {
			const user = await loginWithGoogle();
			if (isSignup) {
				await afterRegist($page.url);
			} else {
				await afterLogin($page.url);
			}
			messagesStore.hide()
		} catch (e: any) {
			if (e.code === 'auth/popup-closed-by-user') {
				return;
			}
			messagesStore.showError()
		}
	};
</script>

<div class="row">
	<div class="col">
		<button on:click={loginGoogle} class="btn btn-primary w-100">Login With Google</button>
	</div>
</div>
