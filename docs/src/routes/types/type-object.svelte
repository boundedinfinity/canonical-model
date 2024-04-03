<script lang="ts">
    import TypeItem from "./type-item.svelte";
</script>

<TypeItem name="object">
    <svelte:fragment slot="props">
        <li>
            <h4>properties</h4>

            <p>The type information for the properties in this object.</p>
        </li>
    </svelte:fragment>
</TypeItem>
