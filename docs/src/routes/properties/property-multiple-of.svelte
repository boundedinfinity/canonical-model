<script lang="ts">
    import Item from "./property-item.svelte";
</script>

<Item name="mulitple-of">
    <p>A description of this field.</p>

    <p>
        The integer constraint must be a multiple of the values. Meaning that if
        the result of the <code>value</code> modulo <code>multiple-of</code>
        evaluates to <code>0</code> will satisfy this contraint.
    </p>
</Item>
