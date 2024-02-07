<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { logout } from '$lib/firebase/auth.client';
	import messagesStore from '$lib/stores/messages.store';

	export let isLoggedIn: boolean;

	let isOpen = false;
	const toggleMenu = () => {
		isOpen = !isOpen;
	};

	const onLogout = async () => {
		try {
            await logout()
			console.log("goto")
            goto('/')
		} catch (e) {
			// console.log(e);
            messagesStore.showError()
		}
	};
</script>

<nav class="navbar navbar-expand-lg">
	<div class="container-fluid">
		<a class="navbar-brand" href="/">TaskTree</a>
		<button
			class="navbar-toggler"
			type="button"
			data-bs-toggle="collapse"
			data-bs-target="#navbarNav"
			aria-controls="navbarNav"
			aria-expanded="false"
			aria-label="Toggle navigation"
			on:click={toggleMenu}
		>
			<span class="navbar-toggler-icon" />
		</button>
		<div class:show={isOpen} class="collapse navbar-collapse" id="navbarNav">
			<ul class="navbar-nav">
				{#if isLoggedIn}
					<li class="nav-item">
						<a
							class:active={$page.url.pathname === '/'}
							class="nav-link"
							aria-current="page"
							href="/">Home</a
						>
					</li>
					<li class="nav-item">
						<a class:active={$page.url.pathname === '/about'} class="nav-link" href="/about"
							>About</a
						>
					</li>
					<li class="nav-item">
						<a class:active={$page.url.pathname === '/base-task'} class="nav-link" href="/base-task"
							>Base Task</a
						>
					</li>
					<li class="nav-item">
						<a class:active={$page.url.pathname === '/account'} class="nav-link" href="/account"
							>Account</a
						>
					</li>
					<li class="nav-item">
						<!-- svelte-ignore a11y-click-events-have-key-events -->
						<!-- svelte-ignore a11y-no-static-element-interactions -->
						<span on:click={onLogout} class="nav-link">Logout</span>
					</li>
				{:else}
					<li class="nav-item">
						<a
							class:active={$page.url.pathname === '/'}
							class="nav-link"
							aria-current="page"
							href="/">Home</a
						>
					</li>
					<li class="nav-item">
						<a class:active={$page.url.pathname === '/about'} class="nav-link" href="/about"
							>About</a
						>
					</li>

					<li class="nav-item">
						<a class:active={$page.url.pathname === '/login'} class="nav-link" href="/login"
							>Login</a
						>
					</li>
					<li class="nav-item">
						<a class:active={$page.url.pathname === '/signup'} class="nav-link" href="/signup"
							>Sign Up</a
						>
					</li>
				{/if}
			</ul>
		</div>
	</div>
</nav>

<style>
	.nav-item > span {
		cursor: pointer;
	}
</style>
