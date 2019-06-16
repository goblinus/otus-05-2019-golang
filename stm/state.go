package stm

import(
	"fmt"
	"strconv"
)

const EmptyStringValue = ""
const FirstSymbol 	= 0
const ErrorCode 	= 1
const EqualValues 	= 0

const StartProcessState = 1
const TerminatedState 	= 5

const DigitSymbol 		= 1
const PrintableSymbol  	= 2
const BackSlashSymbol 	= 3

type State struct {
	state int
	source string
	sourceLength int
	currentSymbol string
	currentIndex int
}

func (s *State) Init(source string) {
	s.source = source
	s.currentIndex = FirstSymbol
	s.sourceLength = len(s.source)
	s.setState(StartProcessState)
}

func (s *State) CurrentSymbol() string {
	return string(s.source[s.currentIndex])
}

func (s *State) Next() bool {
	var result string

	//Если конец строки прекращаем процесс перебора
	if s.currentIndex >= s.sourceLength {
		s.setState(TerminatedState)
		return true
	}

	
	switch symbolType := getSymbolType(s.CurrentSymbol()); symbolType {
	
	case DigitSymbol: 		//Если текущим символом является цифра, значит ошибка
		s.setState(TerminatedState)
		return false

	case BackSlashSymbol:	
		//Если обратный слэш является последним символом последовательности, ошибка:
		//одиночный слэш недопустим по правилам
		s.currentIndex++
		if s.currentIndex >= s.sourceLength {
			s.setState(TerminatedState)
			return false
		}
		
		result := s.CurrentSymbol()
		//Проверка: печатный символ - последний в последовательности
		s.currentIndex++
		if s.currentIndex >= s.sourceLength {
			s.setState(TerminatedState)
		} 
		
		if TerminatedState != s.state  {
			checkedSymbol := s.CurrentSymbol()
			if DigitSymbol == getSymbolType(checkedSymbol) {
				limit, _ := strconv.Atoi(checkedSymbol)
				result = repeat(result, limit)
				s.currentIndex++ //Переходим на следующий необработанный символ
			}
		}
		fmt.Print(result)
	
	default:
		result = s.CurrentSymbol()
		
		s.currentIndex++
		if s.currentIndex >= s.sourceLength {
			s.setState(TerminatedState)
		}

		if s.state != TerminatedState {
			checkedSymbol := s.CurrentSymbol()
			if DigitSymbol == getSymbolType(checkedSymbol) {
				limit, _ := strconv.Atoi(checkedSymbol)
				result = repeat(result, limit)
				s.currentIndex++ //Переходим на следующий необработанный символ
			}
		}
		fmt.Print(result)
	}

	return true
}

func (s *State) CurrentState() int {
	return s.state
}

func (s *State) Rewind() {
	s.currentIndex = FirstSymbol
	s.setState(StartProcessState)
}

//Установка текущего состояния с проверкой допустимых значений
func (s *State) setState(newState int) bool {
	relevantValues := map[int]bool{
		StartProcessState:true,
		TerminatedState:true,
	}
	
	if false == relevantValues[newState] { 
		return false 
	}
	
	s.state = newState
	return true
}