inp="rw.csv"
set datafile separator "," 
set terminal wxt 1 font "Verdana,12" size 1300,700 persist
set xrange [*:*]
set yrange [*:*]
set grid xtics ytics
set style fill solid
unset bars # normal
unset boxwidth
unset mxtics
set xtics out nomirror
set ytics nomirror
plot inp using 1:2 with lines
