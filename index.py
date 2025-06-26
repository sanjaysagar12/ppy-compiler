print('<!DOCTYPE html>')
print('<html lang="en">')
print('<head>')
print('    <meta charset="UTF-8">')
print('    <meta name="viewport" content="width=device-width, initial-scale=1.0">')
print('    <title>Multiplication\' Table</title>')
print('    <style>')
print('        table {')
print('            border-collapse: collapse;')
print('            width: 100%;')
print('            max-width: 800px;')
print('            margin: 0 auto;')
print('        }')
print('        th, td {')
print('            border: 1px solid black;')
print('            padding: 8px;')
print('            text-align: center;')
print('        }')
print('        th {')
print('            background-color: #f2f2f2;')
print('        }')
print('        h1 {')
print('            text-align: center;')
print('        }')
print('    </style>')
print('</head>')
print('<body>')
print('    <h1>Multiplication Table</h1>')
print('    <table>')
print('        <tr>')
print('            <th>×</th>')
# Print table headers (1-10)
for i in range(1, 11):
    print('                <th>')
    print( i)
    print('</th>')
# end of headers loop
print('        </tr>')
# Create rows for the multiplication table
for row in range(1, 11):
    print('            <tr>')
    print('                <th>' + str(row) + '</th>')
    # Create cells for each row
    for col in range(1, 11):
        # Calculate the product
        product = row * col
        print('                    <td>' + str(product) + '</td>')
    # end of column loop
    print('            </tr>')
# end of row loop
print('    </table>')
print('    <p>Generated on: 2025-06-26 15:03:49 by sanjaysagar12</p>')
print('</body>')
print('</html>')
