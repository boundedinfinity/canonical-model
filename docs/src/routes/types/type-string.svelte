<script lang="ts">
    import TypeCommon from "./type-common.svelte";
    import Item from "./property-item.svelte";
    
    export let typeId = "string"
</script>

<TypeCommon {typeId} builtIn>
    <svelte:fragment slot="props">
        <li>
            <Item name="min">
                <p>The inclusive minimum constraint value of the {typeId}.</p>
            </Item>
        </li>
        <li>
            <Item name="max">
                <p>The inclusive maximum constraint value of the {typeId}.</p>
            </Item>
        </li>
        <li>
            <Item name="regex">
                <p>Regular expression for the {typeId} type.</p>
            </Item>
        </li>
        <li>
            <Item name="abnf">
                <p>The ABNF rules contraint for the {typeId} type.</p>
            </Item>
        </li>
    </svelte:fragment>
</TypeCommon>
