package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type UserRole struct {
	ID     string
	UserId string
	Role   string
}

// User struct to hold user data
type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"` // Changed from Name to Username
	Email    string    `json:"email"`
	Password string    `json:"password"` // New field for password
	Value1   string    `json:"value1"`   // Example for additional values
	Value2   string    `json:"value2"`
	Value3   string    `json:"value3"`
	Value4   string    `json:"value4"`
	Value5   string    `json:"value5"`
	Value6   string    `json:"value6"`
	Value7   string    `json:"value7"`
	Value8   string    `json:"value8"`
	Value9   string    `json:"value9"`
	Value10  string    `json:"value10"`
	Value11  string    `json:"value11"`
	Value12  string    `json:"value12"`
	Value13  string    `json:"value13"`
	Value14  string    `json:"value14"`
	Value15  string    `json:"value15"`
	Value16  string    `json:"value16"`
	Value17  string    `json:"value17"`
	Value18  string    `json:"value18"`
	Value19  string    `json:"value19"`
	Value20  string    `json:"value20"`
	Value21  string    `json:"value21"`
	Value22  string    `json:"value22"`
	Value23  string    `json:"value23"`
	Value24  string    `json:"value24"`
	Value25  string    `json:"value25"`
	Value26  string    `json:"value26"`
	Value27  string    `json:"value27"`
	Value28  string    `json:"value28"`
	Value29  string    `json:"value29"`
	Value30  string    `json:"value30"`
	Value31  string    `json:"value31"`
	Value32  string    `json:"value32"`
	Value33  string    `json:"value33"`
	Value34  string    `json:"value34"`
	Value35  string    `json:"value35"`
	Value36  string    `json:"value36"`
	Value37  string    `json:"value37"`
	Value38  string    `json:"value38"`
	Value39  string    `json:"value39"`
	Value40  string    `json:"value40"`
	Role     *UserRole `json:"role"`
}

func main() {
	// Cluster configuration
	cluster := gocql.NewCluster("127.0.0.1:9042") // Ensure correct port
	cluster.Keyspace = "my_saas"                  // Your keyspace name

	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println("Error connecting to Cassandra:", err)
		return
	}
	defer session.Close()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		var (
			id       string
			username string
			email    string
			password string
			value1   string
			value2   string
			value3   string
			value4   string
			value5   string
			value6   string
			value7   string
			value8   string
			value9   string
			value10  string
			value11  string
			value12  string
			value13  string
			value14  string
			value15  string
			value16  string
			value17  string
			value18  string
			value19  string
			value20  string
			value21  string
			value22  string
			value23  string
			value24  string
			value25  string
			value26  string
			value27  string
			value28  string
			value29  string
			value30  string
			value31  string
			value32  string
			value33  string
			value34  string
			value35  string
			value36  string
			value37  string
			value38  string
			value39  string
			value40  string
		)

		var users []User

		// Execute the query
		query := session.Query(`SELECT id, username, email, password, value1, value2, value3, value4, value5, value6, 
            value7, value8, value9, value10, value11, value12, value13, value14, value15, value16, value17, value18, 
            value19, value20, value21, value22, value23, value24, value25, value26, value27, value28, value29, 
            value30, value31, value32, value33, value34, value35, value36, value37, value38, value39, value40 
            FROM users`)

		iter := query.Iter()

		// Iterate over results
		for iter.Scan(&id, &username, &email, &password, &value1, &value2, &value3, &value4, &value5, &value6,
			&value7, &value8, &value9, &value10, &value11, &value12, &value13, &value14, &value15, &value16,
			&value17, &value18, &value19, &value20, &value21, &value22, &value23, &value24, &value25, &value26,
			&value27, &value28, &value29, &value30, &value31, &value32, &value33, &value34, &value35, &value36,
			&value37, &value38, &value39, &value40) {
			users = append(users, User{
				ID:       id,
				Username: username,
				Email:    email,
				Password: password, // Include password if needed
				Value1:   value1,
				Value2:   value2,
				Value3:   value3,
				Value4:   value4,
				Value5:   value5,
				Value6:   value6,
				Value7:   value7,
				Value8:   value8,
				Value9:   value9,
				Value10:  value10,
				Value11:  value11,
				Value12:  value12,
				Value13:  value13,
				Value14:  value14,
				Value15:  value15,
				Value16:  value16,
				Value17:  value17,
				Value18:  value18,
				Value19:  value19,
				Value20:  value20,
				Value21:  value21,
				Value22:  value22,
				Value23:  value23,
				Value24:  value24,
				Value25:  value25,
				Value26:  value26,
				Value27:  value27,
				Value28:  value28,
				Value29:  value29,
				Value30:  value30,
				Value31:  value31,
				Value32:  value32,
				Value33:  value33,
				Value34:  value34,
				Value35:  value35,
				Value36:  value36,
				Value37:  value37,
				Value38:  value38,
				Value39:  value39,
				Value40:  value40,
			})
		}

		// Check for errors during iteration
		if err := iter.Close(); err != nil {
			fmt.Println("Error closing iterator:", err) // Log the specific error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error closing iterator"})
			return
		}

		// Return users in JSON format
		c.JSON(http.StatusOK, users)
	})

	router.GET("/users", func(ctx *gin.Context) {
		var (
			id       string
			username string
			email    string
			password string
			value1   string
			value2   string
			value3   string
			value4   string
			value5   string
			value6   string
			value7   string
			value8   string
			value9   string
			value10  string
			value11  string
			value12  string
			value13  string
			value14  string
			value15  string
			value16  string
			value17  string
			value18  string
			value19  string
			value20  string
			value21  string
			value22  string
			value23  string
			value24  string
			value25  string
			value26  string
			value27  string
			value28  string
			value29  string
			value30  string
			value31  string
			value32  string
			value33  string
			value34  string
			value35  string
			value36  string
			value37  string
			value38  string
			value39  string
			value40  string
		)

		var users []User

		// Execute the query
		query := session.Query(`SELECT id, username, email, password, value1, value2, value3, value4, value5, value6, 
            value7, value8, value9, value10, value11, value12, value13, value14, value15, value16, value17, value18, 
            value19, value20, value21, value22, value23, value24, value25, value26, value27, value28, value29, 
            value30, value31, value32, value33, value34, value35, value36, value37, value38, value39, value40 
            FROM users`)

		iter := query.Iter()

		// Iterate over results
		for iter.Scan(&id, &username, &email, &password, &value1, &value2, &value3, &value4, &value5, &value6,
			&value7, &value8, &value9, &value10, &value11, &value12, &value13, &value14, &value15, &value16,
			&value17, &value18, &value19, &value20, &value21, &value22, &value23, &value24, &value25, &value26,
			&value27, &value28, &value29, &value30, &value31, &value32, &value33, &value34, &value35, &value36,
			&value37, &value38, &value39, &value40) {

			var (
				roleid string
				userId string
				role   string
			)
			var userRole *UserRole = nil
			roleQuery := session.Query(`SELECT id, user_id, role FROM user_role WHERE user_id = ?`, id)
			roleIter := roleQuery.Iter()

			if roleIter.Scan(&roleid, &userId, &role) {
				userRole = &UserRole{
					ID:     roleid,
					UserId: userId,
					Role:   role,
				}
			}

			roleIter.Close()

			users = append(users, User{
				ID:       id,
				Username: username,
				Email:    email,
				Password: password, // Include password if needed
				Value1:   value1,
				Value2:   value2,
				Value3:   value3,
				Value4:   value4,
				Value5:   value5,
				Value6:   value6,
				Value7:   value7,
				Value8:   value8,
				Value9:   value9,
				Value10:  value10,
				Value11:  value11,
				Value12:  value12,
				Value13:  value13,
				Value14:  value14,
				Value15:  value15,
				Value16:  value16,
				Value17:  value17,
				Value18:  value18,
				Value19:  value19,
				Value20:  value20,
				Value21:  value21,
				Value22:  value22,
				Value23:  value23,
				Value24:  value24,
				Value25:  value25,
				Value26:  value26,
				Value27:  value27,
				Value28:  value28,
				Value29:  value29,
				Value30:  value30,
				Value31:  value31,
				Value32:  value32,
				Value33:  value33,
				Value34:  value34,
				Value35:  value35,
				Value36:  value36,
				Value37:  value37,
				Value38:  value38,
				Value39:  value39,
				Value40:  value40,
				Role:     userRole,
			})
		}

		// Check for errors during iteration
		if err := iter.Close(); err != nil {
			fmt.Println("Error closing iterator:", err) // Log the specific error
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error closing iterator"})
			return
		}

		// Return users in JSON format
		ctx.JSON(http.StatusOK, users)
	})
	// Start the server
	router.Run(":8080")
}
