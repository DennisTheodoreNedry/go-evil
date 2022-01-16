for LINE in $(cat examples/examples.txt)
    do 
        echo "----------------"$LINE"----------------"
		echo "# Stage 1: Compiling file"
        ./gevil -tm true -f $LINE -o test_file -exe test
        if [ $? -ge 1 ] 
        then
            echo "* Failed to compile. Hopefully there is some output that might tell what went wrong"
            exit 1
        fi
        echo "* Successfully compiled"
        
        echo "# Stage 2: Running file"
        ./output/test_file.test
        if [ $? -ge 1 ] 
        then
            echo "* Failed to run. Hopefully there is some output that might tell what went wrong"
            exit 1
        fi
        echo "* Successfully run"
        echo -e "--------------------------------\n"
        
        echo "# Stage 3: Cleanup"
        make clean_output
        if [ $? -ge 1 ] 
        then
            echo "* Failed to cleanup. Hopefully there is some output that might tell what went wrong"
            exit 1
        fi
        echo "* Successfully cleaned up"
        echo -e "--------------------------------\n"

	done;

echo "All tests successfully compiled!"