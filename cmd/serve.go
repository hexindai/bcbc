package cmd

import (
	"log"
	"net/http"

	"github.com/hexindai/bcbc/response"
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
	Short: "Serve as a http server",
	Long:  "\nServe as a http server",
	Run: func(cmd *cobra.Command, args []string) {

		http.HandleFunc("/cardInfo.json", func(w http.ResponseWriter, req *http.Request) {
			cardNo := req.FormValue("cardNo")
			cbcr := response.New(cardNo)

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			cbcr.WriteResponse(w, response.JSONContentType)

		})
		log.Printf("Bankinfo server listen on port %s", port)
		log.Fatal(http.ListenAndServe(port, nil))
	},
}
