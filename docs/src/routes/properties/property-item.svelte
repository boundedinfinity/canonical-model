<script lang="ts">
    export let name: string
    export let common:boolean = false
</script>

<header>
    <h4>{name}</h4>
    {#if common}<small>common</small>{/if}
</header>

<slot />

<style>
    header {
        display: flex;
        align-items: baseline;
        gap: 1ch;
        

        & h4 {
            color: var(--lime-6);
        }

        & small {
            color: var(--lime-4);
        }
    }
</style>
