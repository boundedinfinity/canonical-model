<script lang="ts">
    import TypeCommon from './type-common.svelte'
    import Item from "./property-item.svelte";

    export let typeId: string
</script>

<TypeCommon {typeId} builtIn>
    <svelte:fragment slot="props">
        <li>
            <Item name="min">
                <p>The inclusive minimum constraint value of the {typeId}.</p>
            </Item>
        </li>
        <li>
            <Item name="exclusive-min">
                <p>The exclusive minimum constraint value of the {typeId}.</p>
                <p>This is mutually execlusive with the <code>min</code> constraint.</p>
            </Item>
        </li>
        <li>
            <Item name="max">
                <p>The inclusive maximum constraint value of the {typeId}.</p>
            </Item>
        </li>
        <li>
            <Item name="exclusive-max">
                <p>The exclusive maximum constraint value of the {typeId}.</p>
                <p>This is mutually execlusive with the <code>maximum</code> constraint.</p>
            </Item>
        </li>
        <li>
            <Item name="mulitple-of">
                <p>A description of this field.</p>
            
                <p>
                    The <code>{typeId}</code> constraint must be a multiple of this value. 
                </p>
                <p>
                    Any value that satifies the following formula:
                </p>
                <div class="multiple-of">
                   <code>value</code> modulo <code>multiple-of</code> == <code>0</code>
                </div>
            </Item>
            
        </li>
        <li>
            <Item name="invalid-list">
                <p>A list of valid values for the {typeId}.</p>
                <p>If this property is defined, then any values not listed here are accepted.</p>
            </Item>
        </li>
        <li>
            <Item name="valid-list">
                <p>A list of valid values for the {typeId}.</p>
                <p>If this property is defined, then only values listed here are accepted.</p>
            </Item>
            
        </li>
    </svelte:fragment>
</TypeCommon>

<style>
    .multiple-of {
        padding-left: 2rem;
    }
</style>
