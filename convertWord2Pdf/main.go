package main

import(
	"os/exec"
)

func main(){
	arg0 := "lowriter"
	arg1 := "--invisible" //This command is optional, it will help to disable the splash screen of LibreOffice.
	arg2 := "--convert-to"
	arg3 := "pdf:writer_pdf_Export"
	path := "/Formulário_de_acompanhamento_de_trabalho_de_Iniciação_Científica_28-02-24"
	nout, err := exec.Command(arg0,arg1,arg2,arg3,path).Output()
}