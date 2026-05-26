package sql

type sqlDatabase struct {
	Tables  []*sqlTable
	Indexes []*sqlIndex
}

func (this *sqlDatabase) Init() error {
	for _, table := range this.Tables {
		if err := table.Init(this); err != nil {
			return err
		}
	}

	return nil
}

func (this *sqlDatabase) AddTable(table *sqlTable) error {
	// TODO check for duplicate table name and column names
	// TODO check or valid table name and column names
	table.Database = this
	this.Tables = append(this.Tables, table)
	return nil
}

func (this *sqlDatabase) GetTable(name string) (*sqlTable, bool) {
	for _, table := range this.Tables {
		if table.Name == name {
			return table, true
		}
	}

	return nil, false
}

func (this *sqlDatabase) GetIndex(name string) (*sqlIndex, bool) {
	for _, index := range this.Indexes {
		if index.Name == name {
			return index, true
		}
	}

	return nil, false
}
