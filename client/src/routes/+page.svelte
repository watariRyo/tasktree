<script lang="ts">
    import { goto } from "$app/navigation"
	import authStore from "$lib/stores/auth.store";

    export let data;
    // export let pageData = $page.data;

    let isLoggedIn = false;
    $: isLoggedIn = $authStore.isActive ? $authStore.isLoggedIn : data.isLoggedIn;

    const handleUpdateBaseTask = async() => {
        await goto('/base-task')
    }
</script>

{#if isLoggedIn}
<div class="containerr">
    <div class="row mt-2">
        <div class="col-2">
            <button type="button" class="btn btn-primary w-100" on:click={handleUpdateBaseTask}>Add / Edit Base Task</button>
        </div>
    </div>
</div>

    
{:else}
    <h1>isLoggedOff</h1>
{/if}
