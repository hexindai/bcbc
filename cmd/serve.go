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
	serveCmd.Flags().StringVarP(&port, "port", "p", ":3232", "bcbc http server listening on port")
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve as a http server",
	Long:  "Serve as a http server listening on port 3232 by default",
	Run: func(cmd *cobra.Command, args []string) {

		http.HandleFunc("/cardInfo.json", func(w http.ResponseWriter, req *http.Request) {

			cardNo := req.FormValue("cardNo")
			cbcr := response.New(cardNo)

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("Server", "bcbc/"+BCBCVERSION)
			cbcr.WriteResponse(w, response.JSONContentType)

		})
		log.Printf("bcbc server listening on port %s", port)
		log.Fatal(http.ListenAndServe(port, nil))
	},
}
