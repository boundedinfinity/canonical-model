<script lang="ts">
    import TypeCommon from "./type-common.svelte";
    import Item from "./property-item.svelte";
</script>

<TypeCommon typeId="array" builtIn>
    <svelte:fragment slot="props">
        <li>
            <Item name="items">
                <p>The type information for the items in this arrays.</p>
            </Item>
        </li>
    </svelte:fragment>
</TypeCommon>
