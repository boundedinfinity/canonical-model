<script lang="ts">
    import Item from "./property-item.svelte";
    export let typeId: string;
</script>

<Item name="type-id">
    <p>
        The <code>type-id</code> for this entity is <code>{typeId}</code>.
    </p>
</Item>
