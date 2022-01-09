for LINE in $(cat examples/examples.txt)
    do 
        echo "---------------------------------"
		echo "Compiling file" $LINE
        ./gevil -f $LINE
        if [ $? -ge 1 ] 
        then
            echo "* Failed to compile. Hopefully there is some output that might tell what went wrong"
            exit 1
        fi
        echo "* Successfully compiled"
        echo -e "---------------------------------\n"

	done;