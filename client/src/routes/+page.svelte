<script lang="ts">
	import { sendJWTToken } from "$lib/firebase/auth.client.js";
	import authStore from "$lib/stores/auth.store";

    export let data;

    let isLoggedIn = false;
    $: isLoggedIn = $authStore.isActive ? $authStore.isLoggedIn : data.isLoggedIn;
</script>

{#if isLoggedIn}
    <h1>isLoggedIn</h1>
{:else}
    <h1>isLoggedOff</h1>
{/if}
