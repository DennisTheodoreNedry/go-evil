/*MIT License

Copyright (c) 2020 --///--//--

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

#ifndef ARGS_HPP
#define ARGS_HPP

#include <string>
#include <iostream>
#include <vector>

struct argument{
  std::string lName = ""; // The long version of the argument, i.e. --encrypt
  bool set = false; // Has this argument been called?
  std::string sName = ""; // The short version of the argument, i.e. -e
  std::string helpMsg = ""; // When help is called, what should help print for information about this function?
  bool needValue = false; // Those argument need a value to function?
  std::string value = ""; // The value that which the argument holds, used if you want to get whatever the used printed
};

class arg{
private:
  std::vector<argument*> internalArg; // This holds all our arguments and their values
public:
  void addArgument(std::string, std::string, bool, std::string); // For when you have a longer version and a smaller
  void addArgument(std::string, bool, std::string); // Only for smaller versions
  void parseArgs(unsigned int, char**); // Checks which arguments that has actually been typed and raises an error if a undeclared one is found
  bool checkArg(std::string); // Checks if the passed argument has been set or not, returns true if thats the case
  std::string getValue(std::string); // Prints the contents of the passed argument
  void help(); // Prints the help screen
};

void arg::addArgument(std::string name, std::string sName, bool needValue, std::string helpMsg){
  argument *newArg = new argument();
  newArg->lName = name;
  newArg->sName = sName;
  newArg->helpMsg = helpMsg;
  newArg->needValue = needValue;
  this->internalArg.push_back(newArg);
}

void arg::addArgument(std::string name, bool needValue, std::string helpMsg){
  this->addArgument(name, "None", needValue, helpMsg);
}

void arg::help(){
  std::cout << "#### Definied Arguments ####" << std::endl;
  for (argument* arg:this->internalArg){
    std::cout << arg->lName << ((arg->sName != "None")? (", "+ arg->sName):"") << " ";
    std::cout << ((arg->needValue)? "<value>":"") << " | " << arg->helpMsg << std::endl;
  }
}

void arg::parseArgs(unsigned int argc, char** argv){
  bool foundArg = false;
  for (unsigned int i = 1; i < argc; i++){
    foundArg = false;
    for (argument* arg:this->internalArg){
      if (argv[i] == arg->lName || argv[i] == arg->sName){
        arg->set = true;
        if (arg->needValue){
          if (!argv[i+1]){ std::cout << "Error: The argument "<<argv[i]<<" needs an argument to work!" << std::endl;}
          else{ arg->value = argv[i+1]; i++; } // i++ so that we dont generate an error message with the actual value to this argument
        }
        foundArg = true;
        break;
      }
    }
    if (!foundArg){ std::cout << "Error: Undeclared argument "<< argv[i] << std::endl; }
  }
}

bool arg::checkArg(std::string varName){
  bool toReturn = false;
  for (argument* arg:this->internalArg){
    if ((arg->lName == varName || arg->sName == varName) && arg->set){
      toReturn = true;
      break;
    }else{ toReturn = false; } // To create less bugs (i hope)
  }
  return toReturn;
}

std::string arg::getValue(std::string argName){
  std::string toReturn = "";
  for (argument* arg:this->internalArg){
    if ((arg->lName == argName || arg->sName == argName) && arg->set){
      toReturn = arg->value;
      break;
    }
  }
  return toReturn;
}
#endif
