NAME	:= ft_cat
SRCS	:= main.go

all: $(NAME)

$(NAME): $(SRCS)
	go build -o $(NAME) $(SRCS)

clean:
	rm -rf $(NAME)

fclean: clean

re: fclean all

.PHONY: all clean fclean re