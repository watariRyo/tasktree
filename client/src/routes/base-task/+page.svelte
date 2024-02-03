<script lang="ts">
	import { client } from "$lib/helpers/api.helper";
	import { onMount } from "svelte";
    import { PUBLIC_BACKEND_FQDN } from '$env/static/public';
    import { ApiError } from '../../models/api.error';
    

    onMount(async() => {
        try {
            const res = await client(`${PUBLIC_BACKEND_FQDN}/api/v1/base`, {
                method: 'GET',
                headers: {
                    Origin: location.origin
                }
            })
        } catch (e: any) {
            console.log(e.message)
        }
    })

</script>

<h1>Base Task Schedule</h1>