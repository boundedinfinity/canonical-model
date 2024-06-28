<script lang="ts">
    import Link from './property-link.svelte'
    import Item from './property-item.svelte'

    export let typeId: string;
    export let builtIn: boolean = false;
    export let expanded = false;
</script>

<section>
    <header>
        {#if expanded}
            <button on:click={() => (expanded = false)}>^</button>
        {:else}
            <button on:click={() => (expanded = true)}>&gt;</button>
        {/if}

        <h3>{typeId}</h3>
        {#if builtIn}<small>built-in</small>{/if}
    </header>

    {#if expanded}
        {#if $$slots["desc"]}
            <div class="desc">
                <slot name="desc" />
            </div>
        {/if}

        <ul>
            <li>
                <Item name="type-id" common>
                    <p>
                        The <code>type-id</code> for this entity is
                        <code>{typeId}</code>.
                    </p>
                </Item>
            </li>
            <li>
                <Item name="name" common>
                    <p>The name of the <code>{typeId}</code> type.</p>
                    <p>
                        Defaults to a language friendly form of the
                        <code>type-id</code> property.
                    </p>
                </Item>
            </li>
            <li>
                <Item name="description" common>
                    <p>The description of the <code>{typeId}</code> type.</p>
                    <p>Inclusion rules:</p>
                    <ul>
                        <li>
                            If this property's description is present it will be
                            included.
                        </li>
                        <li>
                            If this type <code>inherit</code>ed another type,
                            the type's description will be included.
                        </li>
                    </ul>
                </Item>
            </li>
            <li>
                <Item name="required" common>
                    <p>
                        Determines if the <code>{typeId}</code> is <code>required</code>. 
                    </p>
                    <p>
                        If <code>true</code> the value must be present and conform to any of the
                        constraint values.
                    </p>
                    <p>
                        If <code>false</code> the value is optional and conform to any of the
                        constraint values only if present.
                    </p>
                </Item>
            </li>
            <li>
                <Item name="default" common>
                    <p>
                        The default value of the <code>{name}</code> type. If the the default value is defined it must conform to the 
                        contraints defined by the <code>{name}</code> type.
                    </p>
                </Item>
            </li>
            <li><Link /></li>

            <slot name="props" />
        </ul>
    {/if}
</section>

<style>
    header {
        display: flex;
        align-items: baseline;
        gap: 1ch;
    }

    li {
        padding-block: 0.25rem;
    }

    .desc {
        margin-block: 1rem;
        margin-inline: 1rem;
        padding: 1rem;
        border-color: lightblue;
        border-style: solid;
        border-width: 1px;
    }
</style>
