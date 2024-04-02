<script lang="ts">
    import {
        Name,
        Description,
        Required,
        Default,
        Inherit,
        Link,
        TypeId,
    } from "../properties";

    export let name: string
</script>

<section>
    <header>
        <h3>{name}</h3>
    </header>

    <ul>
        <li><TypeId {name}/></li>
        <li><Name /></li>
        <li><Description /></li>
        <li><Required /></li>
        <li><Default /></li>
        <li><Inherit /></li>
        <li><Link /></li>

        <slot name="props" />
    </ul>
</section>
