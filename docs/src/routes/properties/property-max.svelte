<script lang="ts">
    import Item from "./property-item.svelte";

    export let typeId: string;
</script>

<Item name="max">
    <p>The inclusive maximum constraint value of the {typeId}.</p>
</Item>
