function ASSERT() {
	res="$1"
	exp="$2"
	if [ "$res" != "$exp" ]
	then
		echo "Assert failed : $1 should be $2"
		exit 1
	fi
}

echo "test: START"

ASSERT "$(../ft_cat test1.txt)" "$(cat test1.txt)"
ASSERT "$(../ft_cat test2.txt)" "$(cat test2.txt)"
ASSERT "$(../ft_cat test1.txt test2.txt)" "$(cat test1.txt test2.txt)"

echo "test: OK"
