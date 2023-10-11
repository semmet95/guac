// Code generated by ent, DO NOT EDIT.

package hasmetadata

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLTE(FieldID, id))
}

// SourceID applies equality check predicate on the "source_id" field. It's identical to SourceIDEQ.
func SourceID(v int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldSourceID, v))
}

// PackageVersionID applies equality check predicate on the "package_version_id" field. It's identical to PackageVersionIDEQ.
func PackageVersionID(v int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldPackageVersionID, v))
}

// PackageNameID applies equality check predicate on the "package_name_id" field. It's identical to PackageNameIDEQ.
func PackageNameID(v int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldPackageNameID, v))
}

// ArtifactID applies equality check predicate on the "artifact_id" field. It's identical to ArtifactIDEQ.
func ArtifactID(v int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldArtifactID, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldTimestamp, v))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldKey, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldValue, v))
}

// Justification applies equality check predicate on the "justification" field. It's identical to JustificationEQ.
func Justification(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldJustification, v))
}

// Origin applies equality check predicate on the "origin" field. It's identical to OriginEQ.
func Origin(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldOrigin, v))
}

// Collector applies equality check predicate on the "collector" field. It's identical to CollectorEQ.
func Collector(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldCollector, v))
}

// SourceIDEQ applies the EQ predicate on the "source_id" field.
func SourceIDEQ(v int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldSourceID, v))
}

// SourceIDNEQ applies the NEQ predicate on the "source_id" field.
func SourceIDNEQ(v int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNEQ(FieldSourceID, v))
}

// SourceIDIn applies the In predicate on the "source_id" field.
func SourceIDIn(vs ...int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIn(FieldSourceID, vs...))
}

// SourceIDNotIn applies the NotIn predicate on the "source_id" field.
func SourceIDNotIn(vs ...int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotIn(FieldSourceID, vs...))
}

// SourceIDIsNil applies the IsNil predicate on the "source_id" field.
func SourceIDIsNil() predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIsNull(FieldSourceID))
}

// SourceIDNotNil applies the NotNil predicate on the "source_id" field.
func SourceIDNotNil() predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotNull(FieldSourceID))
}

// PackageVersionIDEQ applies the EQ predicate on the "package_version_id" field.
func PackageVersionIDEQ(v int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldPackageVersionID, v))
}

// PackageVersionIDNEQ applies the NEQ predicate on the "package_version_id" field.
func PackageVersionIDNEQ(v int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNEQ(FieldPackageVersionID, v))
}

// PackageVersionIDIn applies the In predicate on the "package_version_id" field.
func PackageVersionIDIn(vs ...int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIn(FieldPackageVersionID, vs...))
}

// PackageVersionIDNotIn applies the NotIn predicate on the "package_version_id" field.
func PackageVersionIDNotIn(vs ...int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotIn(FieldPackageVersionID, vs...))
}

// PackageVersionIDIsNil applies the IsNil predicate on the "package_version_id" field.
func PackageVersionIDIsNil() predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIsNull(FieldPackageVersionID))
}

// PackageVersionIDNotNil applies the NotNil predicate on the "package_version_id" field.
func PackageVersionIDNotNil() predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotNull(FieldPackageVersionID))
}

// PackageNameIDEQ applies the EQ predicate on the "package_name_id" field.
func PackageNameIDEQ(v int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldPackageNameID, v))
}

// PackageNameIDNEQ applies the NEQ predicate on the "package_name_id" field.
func PackageNameIDNEQ(v int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNEQ(FieldPackageNameID, v))
}

// PackageNameIDIn applies the In predicate on the "package_name_id" field.
func PackageNameIDIn(vs ...int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIn(FieldPackageNameID, vs...))
}

// PackageNameIDNotIn applies the NotIn predicate on the "package_name_id" field.
func PackageNameIDNotIn(vs ...int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotIn(FieldPackageNameID, vs...))
}

// PackageNameIDIsNil applies the IsNil predicate on the "package_name_id" field.
func PackageNameIDIsNil() predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIsNull(FieldPackageNameID))
}

// PackageNameIDNotNil applies the NotNil predicate on the "package_name_id" field.
func PackageNameIDNotNil() predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotNull(FieldPackageNameID))
}

// ArtifactIDEQ applies the EQ predicate on the "artifact_id" field.
func ArtifactIDEQ(v int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldArtifactID, v))
}

// ArtifactIDNEQ applies the NEQ predicate on the "artifact_id" field.
func ArtifactIDNEQ(v int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNEQ(FieldArtifactID, v))
}

// ArtifactIDIn applies the In predicate on the "artifact_id" field.
func ArtifactIDIn(vs ...int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIn(FieldArtifactID, vs...))
}

// ArtifactIDNotIn applies the NotIn predicate on the "artifact_id" field.
func ArtifactIDNotIn(vs ...int) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotIn(FieldArtifactID, vs...))
}

// ArtifactIDIsNil applies the IsNil predicate on the "artifact_id" field.
func ArtifactIDIsNil() predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIsNull(FieldArtifactID))
}

// ArtifactIDNotNil applies the NotNil predicate on the "artifact_id" field.
func ArtifactIDNotNil() predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotNull(FieldArtifactID))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLTE(FieldTimestamp, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldContainsFold(FieldKey, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldContainsFold(FieldValue, v))
}

// JustificationEQ applies the EQ predicate on the "justification" field.
func JustificationEQ(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldJustification, v))
}

// JustificationNEQ applies the NEQ predicate on the "justification" field.
func JustificationNEQ(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNEQ(FieldJustification, v))
}

// JustificationIn applies the In predicate on the "justification" field.
func JustificationIn(vs ...string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIn(FieldJustification, vs...))
}

// JustificationNotIn applies the NotIn predicate on the "justification" field.
func JustificationNotIn(vs ...string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotIn(FieldJustification, vs...))
}

// JustificationGT applies the GT predicate on the "justification" field.
func JustificationGT(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGT(FieldJustification, v))
}

// JustificationGTE applies the GTE predicate on the "justification" field.
func JustificationGTE(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGTE(FieldJustification, v))
}

// JustificationLT applies the LT predicate on the "justification" field.
func JustificationLT(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLT(FieldJustification, v))
}

// JustificationLTE applies the LTE predicate on the "justification" field.
func JustificationLTE(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLTE(FieldJustification, v))
}

// JustificationContains applies the Contains predicate on the "justification" field.
func JustificationContains(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldContains(FieldJustification, v))
}

// JustificationHasPrefix applies the HasPrefix predicate on the "justification" field.
func JustificationHasPrefix(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldHasPrefix(FieldJustification, v))
}

// JustificationHasSuffix applies the HasSuffix predicate on the "justification" field.
func JustificationHasSuffix(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldHasSuffix(FieldJustification, v))
}

// JustificationEqualFold applies the EqualFold predicate on the "justification" field.
func JustificationEqualFold(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEqualFold(FieldJustification, v))
}

// JustificationContainsFold applies the ContainsFold predicate on the "justification" field.
func JustificationContainsFold(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldContainsFold(FieldJustification, v))
}

// OriginEQ applies the EQ predicate on the "origin" field.
func OriginEQ(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldOrigin, v))
}

// OriginNEQ applies the NEQ predicate on the "origin" field.
func OriginNEQ(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNEQ(FieldOrigin, v))
}

// OriginIn applies the In predicate on the "origin" field.
func OriginIn(vs ...string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIn(FieldOrigin, vs...))
}

// OriginNotIn applies the NotIn predicate on the "origin" field.
func OriginNotIn(vs ...string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotIn(FieldOrigin, vs...))
}

// OriginGT applies the GT predicate on the "origin" field.
func OriginGT(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGT(FieldOrigin, v))
}

// OriginGTE applies the GTE predicate on the "origin" field.
func OriginGTE(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGTE(FieldOrigin, v))
}

// OriginLT applies the LT predicate on the "origin" field.
func OriginLT(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLT(FieldOrigin, v))
}

// OriginLTE applies the LTE predicate on the "origin" field.
func OriginLTE(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLTE(FieldOrigin, v))
}

// OriginContains applies the Contains predicate on the "origin" field.
func OriginContains(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldContains(FieldOrigin, v))
}

// OriginHasPrefix applies the HasPrefix predicate on the "origin" field.
func OriginHasPrefix(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldHasPrefix(FieldOrigin, v))
}

// OriginHasSuffix applies the HasSuffix predicate on the "origin" field.
func OriginHasSuffix(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldHasSuffix(FieldOrigin, v))
}

// OriginEqualFold applies the EqualFold predicate on the "origin" field.
func OriginEqualFold(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEqualFold(FieldOrigin, v))
}

// OriginContainsFold applies the ContainsFold predicate on the "origin" field.
func OriginContainsFold(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldContainsFold(FieldOrigin, v))
}

// CollectorEQ applies the EQ predicate on the "collector" field.
func CollectorEQ(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEQ(FieldCollector, v))
}

// CollectorNEQ applies the NEQ predicate on the "collector" field.
func CollectorNEQ(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNEQ(FieldCollector, v))
}

// CollectorIn applies the In predicate on the "collector" field.
func CollectorIn(vs ...string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldIn(FieldCollector, vs...))
}

// CollectorNotIn applies the NotIn predicate on the "collector" field.
func CollectorNotIn(vs ...string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldNotIn(FieldCollector, vs...))
}

// CollectorGT applies the GT predicate on the "collector" field.
func CollectorGT(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGT(FieldCollector, v))
}

// CollectorGTE applies the GTE predicate on the "collector" field.
func CollectorGTE(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldGTE(FieldCollector, v))
}

// CollectorLT applies the LT predicate on the "collector" field.
func CollectorLT(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLT(FieldCollector, v))
}

// CollectorLTE applies the LTE predicate on the "collector" field.
func CollectorLTE(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldLTE(FieldCollector, v))
}

// CollectorContains applies the Contains predicate on the "collector" field.
func CollectorContains(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldContains(FieldCollector, v))
}

// CollectorHasPrefix applies the HasPrefix predicate on the "collector" field.
func CollectorHasPrefix(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldHasPrefix(FieldCollector, v))
}

// CollectorHasSuffix applies the HasSuffix predicate on the "collector" field.
func CollectorHasSuffix(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldHasSuffix(FieldCollector, v))
}

// CollectorEqualFold applies the EqualFold predicate on the "collector" field.
func CollectorEqualFold(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldEqualFold(FieldCollector, v))
}

// CollectorContainsFold applies the ContainsFold predicate on the "collector" field.
func CollectorContainsFold(v string) predicate.HasMetadata {
	return predicate.HasMetadata(sql.FieldContainsFold(FieldCollector, v))
}

// HasSource applies the HasEdge predicate on the "source" edge.
func HasSource() predicate.HasMetadata {
	return predicate.HasMetadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, SourceTable, SourceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSourceWith applies the HasEdge predicate on the "source" edge with a given conditions (other predicates).
func HasSourceWith(preds ...predicate.SourceName) predicate.HasMetadata {
	return predicate.HasMetadata(func(s *sql.Selector) {
		step := newSourceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPackageVersion applies the HasEdge predicate on the "package_version" edge.
func HasPackageVersion() predicate.HasMetadata {
	return predicate.HasMetadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PackageVersionTable, PackageVersionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPackageVersionWith applies the HasEdge predicate on the "package_version" edge with a given conditions (other predicates).
func HasPackageVersionWith(preds ...predicate.PackageVersion) predicate.HasMetadata {
	return predicate.HasMetadata(func(s *sql.Selector) {
		step := newPackageVersionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAllVersions applies the HasEdge predicate on the "all_versions" edge.
func HasAllVersions() predicate.HasMetadata {
	return predicate.HasMetadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AllVersionsTable, AllVersionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAllVersionsWith applies the HasEdge predicate on the "all_versions" edge with a given conditions (other predicates).
func HasAllVersionsWith(preds ...predicate.PackageName) predicate.HasMetadata {
	return predicate.HasMetadata(func(s *sql.Selector) {
		step := newAllVersionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasArtifact applies the HasEdge predicate on the "artifact" edge.
func HasArtifact() predicate.HasMetadata {
	return predicate.HasMetadata(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ArtifactTable, ArtifactColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArtifactWith applies the HasEdge predicate on the "artifact" edge with a given conditions (other predicates).
func HasArtifactWith(preds ...predicate.Artifact) predicate.HasMetadata {
	return predicate.HasMetadata(func(s *sql.Selector) {
		step := newArtifactStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.HasMetadata) predicate.HasMetadata {
	return predicate.HasMetadata(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.HasMetadata) predicate.HasMetadata {
	return predicate.HasMetadata(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.HasMetadata) predicate.HasMetadata {
	return predicate.HasMetadata(sql.NotPredicates(p))
}