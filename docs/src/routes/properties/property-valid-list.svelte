<script lang="ts">
    import Item from "./property-item.svelte";

    export let typeId: string;
</script>

<Item name="valid-list">
    <p>A list of valid values for the {typeId}.</p>
    <p>If this property is defined, then only values listed here are accepted.</p>
</Item>
