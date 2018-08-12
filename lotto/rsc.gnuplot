inp="rsc1.csv"
set datafile separator "," 
set terminal wxt 1 font "Verdana,12" size 700,700 persist
set xrange [*:*]
set yrange [*:*]
unset grid
unset mxtics
unset xzeroaxis
unset yzeroaxis
unset x2zeroaxis
unset y2zeroaxis
unset zzeroaxis
# set view 90,30
splot inp using 1:($2):($3) with lines ls 5 lc 7
