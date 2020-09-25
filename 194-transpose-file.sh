#!/usr/bin/env bash

# Loop over each column in a row, appending it to a new array that represents
# the new set of transposed rows. Save maximum number of fields for the END
# phase. In the END phase, strip leading spaces created by processing, and
# then print each row.
awk '\
{
  for (i = 1; i <= NF; i++) rows[i] = rows[i] " " $i;
  if (NF > maxFields) maxFields = NF
}
END {
  for (i = 1; i <= maxFields; i++) sub(/^  */, "", rows[i]);
  for (i = 1; i <= maxFields; i++) print rows[i]
}\
' file.txt
