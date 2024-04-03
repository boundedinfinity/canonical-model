<script lang="ts">
    import Item from "./property-item.svelte";
</script>

<Item name="link" common>
    <p>
        Link to another type definition via that type's <code>type-id</code>.
    </p>

    <p>
        This allows a type to link another type, in effect, creating a graph of
        items that this type is linked to.
    </p>
</Item>
