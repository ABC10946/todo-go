# todo-go

## Description
todo-go is a command-line tool for managing your to-do list.

## Installation
To install todo-go, follow these steps:

1. Clone the repository:
    ```
    git clone https://github.com/ABC10946/todo-go.git
    ```

2. Navigate to the project directory:
    ```
    cd todo-go
    ```

3. Build the executable:
    ```
    go build
    ```

4. Run the program:
    ```
    ./todo-go
    ```

## Usage
todo-go supports the following commands:

- `create --title <title> --description <description> --completed <completed>`: Adds a new task to the to-do list.
- `delete --id <todo id>`: Removes a task from the to-do list.
- `list`: Displays all tasks in the to-do list.

Example usage:
```
./todo-go create --title "buy something" --description "buying something in market" --completed false
./todo-go list
./todo-go delete --id 2
```

---

## 説明
todo-goは、あなたのタスクリストを管理するためのコマンドラインツールです。

## インストール
todo-goをインストールするには、以下の手順に従ってください：

1. リポジトリをクローンします：
    ```
    git clone https://github.com/ABC10946/todo-go.git
    ```

2. プロジェクトディレクトリに移動します：
    ```
    cd todo-go
    ```

3. 実行可能ファイルをビルドします：
    ```
    go build
    ```

4. プログラムを実行します：
    ```
    ./todo-go
    ```

## 使用方法
todo-goは以下のコマンドをサポートしています：

- `create --title <title> --description <description> --completed <completed>`: 新しいタスクをタスクリストに追加します。
- `delete --id <todo id>`: タスクをタスクリストから削除します。
- `list`: タスクリスト内のすべてのタスクを表示します。

使用例：
```
./todo-go create --title "買い物" --description "市場で何かを買う" --completed false
./todo-go list
./todo-go delete --id 2
```