package cache


type Series struct {
  Title string
  Seasons []Season
  Description string
  Poster string
}

type Season struct {
  Series Series
  


}

type Movie struct {
  Title string
  Description string



}


type Episode struct {
  Title string
  Description string
  File string



}
