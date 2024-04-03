<script lang="ts">
    import TypeItem from "./type-item.svelte";
</script>

<TypeItem name="array" builtIn>
    <svelte:fragment slot="props">
        <li>
            <h4>items</h4>

            <p>The type information for the items in this arrays.</p>
        </li>
    </svelte:fragment>
</TypeItem>
