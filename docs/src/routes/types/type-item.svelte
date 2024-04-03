<script lang="ts">
    import {
        Name,
        Description,
        Required,
        Default,
        Inherit,
        Link,
        TypeId,
        Type,
    } from "../properties";

    export let name: string
    export let builtIn: boolean = false
</script>

<section>
    <header>
        <h3>{name}</h3>
        {#if builtIn}<small>built-in</small>{/if}
    </header>

    <ul>
        <li><Type {name}/></li>
        <li><TypeId /></li>
        <li><Name /></li>
        <li><Description /></li>
        <li><Required /></li>
        <li><Default /></li>
        <li><Inherit /></li>
        <li><Link /></li>

        <slot name="props" />
    </ul>
</section>

<style>
    header {
        display: flex;
        align-items: baseline;
        gap: 1ch;
    }
</style>
