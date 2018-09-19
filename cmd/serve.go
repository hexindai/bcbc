package cmd

import (
	"log"
	"net/http"

	"github.com/runrioter/bcbc/bank"
	"github.com/spf13/cobra"
)

var (
	port string
)

func init() {
	bcbc.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&port, "port", "p", ":3232", "Bankinfo http server listening port")
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a http bankcard info server",
	Long:  "\nStart http server for searching bankcard info",
	Run: func(cmd *cobra.Command, args []string) {

		http.HandleFunc("/cardInfo.json", func(w http.ResponseWriter, req *http.Request) {
			cardNo := req.FormValue("cardNo")
			cbcr := bank.FetchCardBinCheckByCard(cardNo)

			cbcr.WriteResponse(w, bank.JSONContentType)

		})
		log.Printf("Bankinfo server listen on port %s", port)
		log.Fatal(http.ListenAndServe(port, nil))
	},
}
