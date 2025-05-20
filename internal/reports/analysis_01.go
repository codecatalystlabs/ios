package reports

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
)

func GenerateHTMLSummary(c context.Context, db *sql.DB, table, column string, labelMap map[string]string) (string, error) {
	query := fmt.Sprintf(`SELECT %s, COUNT(*) as count FROM %s GROUP BY %s`, column, table, column)

	rows, err := db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var total int
	type Summary struct {
		Value string
		Count int
	}
	var summaries []Summary

	for rows.Next() {
		var value sql.NullString
		var count int
		if err := rows.Scan(&value, &count); err != nil {
			return "", err
		}
		if value.Valid {

			if labelMap != nil {
				if label, exists := labelMap[value.String]; exists {
					value.String = label
				}
			}

			summaries = append(summaries, Summary{Value: value.String, Count: count})
			total += count
		}
	}

	// HTML with inline CSS for better table styling
	html := `
	<style>
		table {
			width: 100%;
			border-collapse: collapse;
		}
		th, td {
			padding: 2px;
			text-align: left;
		}
		tr:nth-child(even) {
			background-color: #f2f2f2;
		}
		hr {
			border: 1px solid #ccc;
			margin: 10px 0;
		}
	</style>
	<hr>
	<table>
		<tr>
			<th>Value</th>
			<th>Frequency</th>
			<th>Percentage</th>
		</tr>`

	// Add rows for each summary
	for _, summary := range summaries {
		percentage := (float64(summary.Count) / float64(total)) * 100
		html += fmt.Sprintf("<tr><td>%s</td><td>%d</td><td>%.1f%%</td></tr>", summary.Value, summary.Count, percentage)
	}

	// Add totals row at the bottom
	html += fmt.Sprintf(`
	<tr>
		<td><strong>Total</strong></td>
		<td><strong>%d</strong></td>
		<td><strong>100.0%%</strong></td>
	</tr>
	</table>
	<hr>`, total)
	html += "</table>"
	//html += `<hr>`
	return html, nil
}

func GenerateHTMLFrequencySummary(c context.Context, db *sql.DB, table string, fields []string, labelMap map[string]string) (string, error) {
	const MissingLabel = "Missing" // Label for NULL and unmapped values

	// Map to store unique values per field and their frequencies
	frequencies := make(map[string]map[string]int)
	uniqueValues := make(map[string][]string) // Store unique values for column headers

	for _, field := range fields {
		frequencies[field] = make(map[string]int)

		// Query to count occurrences of each unique value in the field
		query := fmt.Sprintf(`SELECT %s, COUNT(*) FROM %s GROUP BY %s ORDER BY %s`, field, table, field, field)

		rows, err := db.Query(query)
		if err != nil {
			return "", err
		}
		defer rows.Close()

		var value sql.NullString
		var count int
		for rows.Next() {
			if err := rows.Scan(&value, &count); err != nil {
				return "", err
			}

			var val string
			if value.Valid {
				val = value.String
			} else {
				val = MissingLabel // Assign Missing label for NULL values
			}

			frequencies[field][val] = count

			if !contains(uniqueValues[field], val) {
				uniqueValues[field] = append(uniqueValues[field], val)
			}
		}
	}

	// Collect all unique values across fields to create column headers
	allUniqueValues := make(map[string]bool)
	for _, values := range uniqueValues {
		for _, val := range values {
			allUniqueValues[val] = true
		}
	}
	allUniqueValues[MissingLabel] = true // Ensure Missing column is always present

	// Convert map keys to sorted slice for consistent column ordering
	sortedUniqueValues := sortedKeys(allUniqueValues)

	// Generate HTML table
	html := `
		<style>
			table {
				width: 100%;
				border-collapse: collapse;
			}
			th, td {
				padding: 8px;
				text-align: left;
				border-bottom: 1px solid #ddd;
			}
			tr:nth-child(even) {
				background-color: #f2f2f2;
			}
			hr {
				border: 1px solid #ccc;
				margin: 10px 0;
			}
		</style>
		<hr>
		<table>
		<tr>
			<th>Variable</th>`

	// Table headers (Unique values, replaced with labels if available)
	for _, val := range sortedUniqueValues {
		label := MissingLabel // Default Missing label
		if val != MissingLabel {
			if mappedLabel, exists := labelMap[val]; exists {
				label = mappedLabel
			} else {
				label = val // Use raw value if no mapping is found
			}
		}
		html += fmt.Sprintf("<th>%s</th>", label)
	}
	html += "</tr>"

	// Populate frequency table
	for _, field := range fields {
		label := field
		if mappedLabel, exists := labelMap[field]; exists {
			label = mappedLabel
		}

		html += fmt.Sprintf("<tr><th>%s</th>", label)

		for _, val := range sortedUniqueValues {
			count := frequencies[field][val]
			html += fmt.Sprintf("<td>%d</td>", count)
		}
		html += "</tr>"
	}

	html += "</table><hr>"

	return html, nil
}

// Checks if a slice contains a string
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Sorts map keys into a slice
func sortedKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys) // Sort alphabetically for consistency
	return keys
}
