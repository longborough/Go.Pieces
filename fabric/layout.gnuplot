inp="new2.csv"
set datafile separator "," 
set grid xtics ytics
set style fill empty
unset bars # normal
unset boxwidth
unset mxtics
set key off
set xtics out 20
set ytics out 10
unset y2tics
set yrange [*:*] 
set xrange [0:425]
plot inp using ($2+0.5*$4):($3+0.5*$5):(0.5*($4-1)):(0.5*($5-1)) with boxxyerrorbars lc 0
set xrange [390:815]
plot inp using ($2+0.5*$4):($3+0.5*$5):(0.5*($4-1)):(0.5*($5-1)) with boxxyerrorbars lc 0
set xrange [800:1225]
plot inp using ($2+0.5*$4):($3+0.5*$5):(0.5*($4-1)):(0.5*($5-1)) with boxxyerrorbars lc 0
