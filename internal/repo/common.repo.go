package repo

import "context"

// TableNameGoCrmUser defines the table name for CRM users.
const TableNameGoCrmUser = "go_crm_user"

// NumberNil is used for comparison to check if rows are affected.
const NumberNil = 0

// ctx is the global background context for operations that don't need cancelation.
var ctx = context.Background()
