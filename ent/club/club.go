// Code generated by entc, DO NOT EDIT.

package club

const (
	// Label holds the string label denoting the club type in the database.
	Label = "club"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSummary holds the string denoting the summary field in the database.
	FieldSummary = "summary"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldHashtags holds the string denoting the hashtags field in the database.
	FieldHashtags = "hashtags"
	// FieldImages holds the string denoting the images field in the database.
	FieldImages = "images"
	// FieldHomepage holds the string denoting the homepage field in the database.
	FieldHomepage = "homepage"
	// FieldInstagram holds the string denoting the instagram field in the database.
	FieldInstagram = "instagram"
	// FieldFacebook holds the string denoting the facebook field in the database.
	FieldFacebook = "facebook"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldRecommended holds the string denoting the recommended field in the database.
	FieldRecommended = "recommended"
	// EdgeLikes holds the string denoting the likes edge name in mutations.
	EdgeLikes = "likes"
	// Table holds the table name of the club in the database.
	Table = "clubs"
	// LikesTable is the table the holds the likes relation/edge.
	LikesTable = "like_clubs"
	// LikesInverseTable is the table name for the LikeClub entity.
	// It exists in this package in order to avoid circular dependency with the "likeclub" package.
	LikesInverseTable = "like_clubs"
	// LikesColumn is the table column denoting the likes relation/edge.
	LikesColumn = "club_likes"
)

// Columns holds all SQL columns for club fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSummary,
	FieldDescription,
	FieldHashtags,
	FieldImages,
	FieldHomepage,
	FieldInstagram,
	FieldFacebook,
	FieldPhone,
	FieldEmail,
	FieldRecommended,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultRecommended holds the default value on creation for the "recommended" field.
	DefaultRecommended bool
)