<script lang="ts">
    import {
        Name,
        Description,
        Required,
        Default,
        Link,
        TypeId,
    } from "../properties";

    export let typeId: string;
    export let builtIn: boolean = false;
    export let expanded = false;
</script>

<section>
    <header>
        {#if expanded}
            <button on:click={() => (expanded = false)}>^</button>
        {:else}
            <button on:click={() => (expanded = true)}>&gt;</button>
        {/if}

        <h3>{typeId}</h3>
        {#if builtIn}<small>built-in</small>{/if}
    </header>

    {#if expanded}
        {#if $$slots["desc"]}
            <div class="desc">
                <slot name="desc" />
            </div>
        {/if}

        <ul>
            <li><TypeId {typeId} /></li>
            <li><Name {typeId} /></li>
            <li><Description {typeId} /></li>
            <li><Required /></li>
            <li><Default name={typeId} /></li>
            <li><Link /></li>

            <slot name="props" />
        </ul>
    {/if}
</section>

<style>
    header {
        display: flex;
        align-items: baseline;
        gap: 1ch;
    }

    li {
        padding-block: 0.25rem;
    }

    .desc {
        margin-block: 1rem;
        margin-inline: 1rem;
        padding: 1rem;
        border-color: lightblue;
        border-style: solid;
        border-width: 1px;
    }
</style>
