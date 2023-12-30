<script lang="ts">
    import Nav from '$lib/components/Nav.svelte';
    import 'bootstrap/dist/css/bootstrap.min.css';
    import '$lib/firebase/firebase.client';
	import authStore from '$lib/stores/auth.store';

    export let data;

    let isLoggedIn = false;
    $: isLoggedIn = $authStore.isActive ? $authStore.isLoggedIn : data.isLoggedIn;
</script>

<Nav isLoggedIn={isLoggedIn}/>
<main class="container">
    <slot />
</main>