<script lang="ts">
    import Item from "./property-item.svelte";
    export let typeId: string
</script>

<Item name="mulitple-of">
    <p>A description of this field.</p>

    <p>
        The <code>{typeId}</code> constraint must be a multiple of this value. 
    </p>
    <p>
        Any value that satifies the following formula:
    </p>
    <div>
       <code>value</code> modulo <code>multiple-of</code> == <code>0</code>
    </div>
</Item>

<style>
    div {
        padding-left: 2rem;
    }
</style>
