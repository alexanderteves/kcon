package main

import (
	"flag"
	"github.com/alexanderteves/kubeassist/pkg/config"
	"log"
	"os"
)

func main() {
	var context, namespace string
	flag.StringVar(&context, "c", "", "Context to switch to (will fail if undefined)")
	flag.StringVar(&namespace, "n", "", "Update default namespace for context")
	flag.Parse()

	kcfgPath := os.Getenv("KUBECONFIG")
	if kcfgPath == "" {
		log.Fatal("KUBECONFIG not set")
	}

	kcfg, err := config.Load(kcfgPath)
	if err != nil {
		log.Fatal(err)
	}

	if context != "" {
		err = kcfg.SetContext(context)
		if err != nil {
			log.Fatal(err)
		}
	}

	if namespace != "" {
		kcfg.SetNamespace(namespace)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = kcfg.Dump(kcfgPath)
	if err != nil {
		log.Fatal(err)
	}
}
