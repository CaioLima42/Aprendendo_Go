package main
import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
// Construtor de MemoryStore
func NewMemoryStore() MemoryStore {
 	return MemoryStore{items: make(map[string]string)}
}
// istancia de memoria dedicada a armazenar as urls {key = hash e value = url}
type MemoryStore struct {
 	items map[string]string
}
// função dedicada a adicionar uma nova url
func (m *MemoryStore) Add(shortendURL, longURL string) error {
	if m.items[shortendURL] != "" {
		return fmt.Errorf("value already exists here")
 	}
	m.items[shortendURL] = longURL
	log.Println(m.items)
	return nil
}
// funcao dedicada a remover uma url
func (m *MemoryStore) Remove(shortenedURL string) error {
 	if m.items[shortenedURL] == "" {
 		return fmt.Errorf("value does not exist here")
 	}
	delete(m.items, shortenedURL)
	return nil
	}
// funcao dedicada a buscar uma url
func (m *MemoryStore) Get(shortendURL string) (string, error) {
	longURL, ok := m.items[shortendURL]
	if !ok {
		return "", fmt.Errorf("no mapped url available here")
	}
	return longURL, nil
	}
func (m *MemoryStore) Show() map[string]string{
	return m.items
}
// interface com todos os metodos que serao usados
type Store interface {
	Add(shortenedURL, longURL string) error
	Remove(shortenedURL string) error
	Get(shortendURL string) (string, error)
	Show() map[string]string
}
// varivavel que salva o dominio e as urls
type AddPath struct {
	domain string
	store Store
}
func (a *AddPath) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// struct deficada a salvar o valor do corpo da requisição
	type addPathRequest struct {
		URL string `json:"url"`
	}
	var parsed addPathRequest
	// verifica se ah algum conteundo na requisição, caso contrario, retorna nil
	// salva a url na variavel addPathRequest
	err := json.NewDecoder(r.Body).Decode(&parsed)
	// envia mensagem de erro caso haja um json sem corpo ou com corpo errado
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("unexpected error :: %v", err)))
		return
	}
	// cria uma nova criptogravia para a url encurtada(sha1)
	h := sha1.New()
	// escreve o hash(toda vez que chamado, acrecenta o novo hash a variavel)
	h.Write([]byte(parsed.URL))
	// Cria de fato o hash
	sum := h.Sum(nil)
	// codifica de hexa para string e pega apenas as 9 primeira casas 
	hash := hex.EncodeToString(sum)[:10]
	// adiciona a url encurtada ao map
	err = a.store.Add(hash, parsed.URL)
	// Verifica se a url ja existe
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("unexpected error :: %v", err)))
		return
 	}
	// variavel dedicada a retornar a url para o usuario
	type addPathResponse struct {
		ShortenedURL string `json:"shortened_url"`
		LongURL string `json:"long_url"`
	}
	// variavel que salva o valor que sera devolvido pelo servidor
	pathResp := addPathResponse{ShortenedURL: fmt.Sprintf("%v/%v", a.domain, hash), LongURL: parsed.URL}
	// indica que as informações seram devolvidas no formato json
	w.Header().Set("Content-Type", "application/json")
	// einvia a resposta que o url foi criado para o cliente
	w.WriteHeader(http.StatusCreated)
	// escreve o json que sera enviado para o cliente
	json.NewEncoder(w).Encode(pathResp)
	}

// varivel responsavel por deletar a url
type DeletePath struct {
	store Store
}
func (p *DeletePath) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// busca o valor hash no json enviado
	hash := mux.Vars(r)["hash"]
	// se o valor for vazio, retorna um erro
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty hash"))
		return
 }
 	// se não for encontrado a url, retorna um erro
	err := p.store.Remove(hash)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("unexpected error :: %v", err)))
		return
 	}
	// escreve o estatus ok e indica que a url foi deletada
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted"))
}
type RedirectPath struct {
 	store Store
}
func (p *RedirectPath) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// busca o valor hash no json enviado
	hash := mux.Vars(r)["hash"]
	// se o valor for vazio, retorna um erro
	if hash == "" {
		log.Println("Flamengooooooo")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty hash"))
		return
	}
	// se não for encontrado a url, retorna um erro
	longURL, err := p.store.Get(hash)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}
	// redireciona o url encurtado para a pagina real
	http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
}

type Allpath struct{
	domain string
	store Store
}

func(a *Allpath) ServeHTTP(w http.ResponseWriter, r *http.Request){
	urls := a.store.Show()
	type FormatedJson struct{
		Hash string `json:"hash"`
		Url string `json:"url"`
	}
	if len(urls) == 0{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Empty List"))
		return
	}
	response := []FormatedJson{}
	for k, v := range urls{
		response = append(response, FormatedJson{Hash: fmt.Sprintf("%v%v", a.domain + "/", k ), Url: v})
	}
	w.Header().Set("Content-Type", "application/json")
	// einvia a resposta que o url foi criado para o cliente
	w.WriteHeader(http.StatusOK)
	// escreve o json que sera enviado para o cliente
	json.NewEncoder(w).Encode(response)
}


type HandleViaStruct struct{}

func (*HandleViaStruct) ServeHTTP(w http.ResponseWriter, r *http.Request){
	log.Print("Hello world received a request.")
	defer log.Print("End hello world request")
	fmt.Fprintf(w, "Hello World via Struct")
}
func main() {
	log.Print("Server is started")
	r := mux.NewRouter()
	redirectPath := "http://localhost:8080/r"
	mem := NewMemoryStore()
	r.Handle("/", &HandleViaStruct{}).Methods("GET")
	r.Handle("/add", &AddPath{domain: redirectPath, store: &mem}).Methods("POST")
	r.Handle("/r/{hash}", &DeletePath{store: &mem}).Methods("DELETE")
	r.Handle("/r/{hash}", &RedirectPath{store: &mem}).Methods("GET")
	r.Handle("/s", &Allpath{domain: redirectPath,store: &mem}).Methods("GET")
	http.ListenAndServe(":8080", r)
}